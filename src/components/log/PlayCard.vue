<template>
  <div>
    <sn-player-btn :player='value.player' size='small'></sn-player-btn>
    {{value.player.user.name}} played {{kind}} card.
  </div>
</template>

<script>
  import Button from '@/components/player/Button'

  export default {
    name: 'sn-log-play-card-msg',
    props: [ 'value' ],
    components: {
      'sn-player-btn': Button
    },
    computed: {
      kind: function () {
        var self = this
        switch (self.value.card.kind) {
          case 'start-lamp':
            return 'lamp'
          case 'start-camel':
            return 'camel'
          default:
            return self.value.card.kind
        }
      }
    }
  }
</script>
