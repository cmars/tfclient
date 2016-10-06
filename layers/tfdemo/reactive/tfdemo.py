import shutil

from charms.reactive import when, when_not, set_state, remove_state
from charms.reactive.helpers import data_changed

from charmhelpers.core.templating import render
from charmhelpers.core import hookenv, host


@when_not('tfdemo.installed')
def install_tfdemo():
    shutil.copyfile("bin/tfwebapp", "/usr/bin/tfwebapp")
    set_state('tfdemo.installed')


@when('tensorflow.available', 'tfdemo.installed')
def tensorflow_available(tf):
    addrs = tf.addrs()
    if data_changed('tfdemo.addr', addrs) and addrs:
        render(source="tfdemo.service",
            target="/etc/systemd/service/tfdemo.service",
            owner="root",
            perms=0o644,
            context={
                'addr': ','.join(addrs),
            })
        if host.service_running('tfdemo'):
            host.service_restart('tfdemo')
        else:
            host.service_start('tfdemo')
    remove_state('tensorflow.available')


@when('website.available')
def website_available(w):
    w.configure(8080)
    hookenv.open_port(8080)
    remove_state('website.available')
