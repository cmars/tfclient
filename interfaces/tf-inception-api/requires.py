import os
import tempfile

from charms.reactive import RelationBase, scopes, hook


class TfInceptionRequires(RelationBase):
    scope = scopes.SERVICE

    @hook('{requires:tf-inception-api}-relation-{joined,changed}')
    def joined(self):
        conv = self._maybe_conversation()
        if conv:
            conv.set_state('{relation_name}.available')

    @hook('{requires:tf-inception-api}-relation-{departed,broken}')
    def departed(self):
        conv = self._maybe_conversation()
        if conv:
            conv.remove_state('{relation_name}.available')

    def _maybe_conversation(self):
        conv = None
        try:
            conv = self.conversation()
        except:
            pass
        return conv

    def _maybe_conversations(self):
        convs = []
        try:
            convs = self.conversations()
        except:
            pass
        return convs

    def addrs(self):
        addrs = set()
        for conv in self.conversations():
            host = conv.get_remote('host')
            port = conv.get_remote('port')
            if host and port:
                addr = '%s:%d' % (host, port)
                addrs.append(addr)
        return addrs
