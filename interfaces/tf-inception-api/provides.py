import json

from charms.reactive import RelationBase, scopes, hook
from charmhelpers.core import hookenv


class TfInceptionProvides(RelationBase):
    scope = scopes.UNIT

    @hook('{provides:tf-inception-api}-relation-{joined,changed}')
    def changed(self):
        conv = self._maybe_conversation()
        if conv:
            self.set_state('{relation_name}.available')

    @hook('{provides:tf-inception-api}-relation-{departed,broken}')
    def departed(self):
        conv = self._maybe_conversation()
        if conv:
            self.remove_state('{relation_name}.available')

    def _maybe_conversation(self):
        conv = None
        try:
            conv = self.conversation()
        except:
            pass
        return conv

    def configure(self, port):
        relation_info = {
            'host': hookenv.unit_get('private-address'),
            'port': port,
        }
        self.set_remote(**relation_info)
