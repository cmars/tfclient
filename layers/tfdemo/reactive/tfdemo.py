import os
import shutil
from subprocess import check_call

from charms.reactive import when, when_not, set_state, remove_state, hook
from charms.reactive.helpers import data_changed

from charmhelpers.core.templating import render
from charmhelpers.core import hookenv, host


@hook('upgrade-charm')
def upgrade_charm():
    install_tfdemo()


@when_not('tfdemo.installed')
def install_tfdemo():
    shutil.copyfile("bin/tfwebapp", "/usr/bin/tfwebapp")
    os.chmod("/usr/bin/tfwebapp", 0o755)
    if host.service_available('tfdemo') and host.service_running('tfdemo'):
        host.service_restart('tfdemo')
    set_state('tfdemo.installed')


@when('client.available', 'tfdemo.installed')
def tensorflow_available(tf):
    addrs = tf.addrs()
    if not addrs:
        if host.service_running('tfdemo'):
            host.service_stop('tfdemo')
        return
    hookenv.open_port(8080)
    ctx = {
        'addr': ','.join(addrs),
    }
    samples_dir = os.path.join(hookenv.charm_dir(), "samples")
    if os.path.exists(samples_dir):
        ctx['samples'] = samples_dir
    render(source="tfdemo.service",
        target="/etc/systemd/system/tfdemo.service",
        owner="root",
        perms=0o644,
        context=ctx,
    )
    check_call(['systemctl', 'daemon-reload'])
    if host.service_running('tfdemo'):
        host.service_restart('tfdemo')
    else:
        host.service_start('tfdemo')
    remove_state('client.available')


@when('website.available')
def website_available(w):
    w.configure(8080)
    remove_state('website.available')
