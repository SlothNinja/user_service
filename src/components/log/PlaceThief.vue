<template>
  <div>
    <sn-player-btn :player='value.player' size='small'></sn-player-btn>
    {{value.player.user.name}} placed thief on {{value.area.card.kind}} at {{row}}{{value.area.column}}.
  </div>
</template>

<script>
  import Text from '@/components/mixins/Text'
  import Button from '@/components/player/Button'

  export default {
    mixins: [ Text ],
    name: 'sn-log-place-thief-msg',
    components: {
      'sn-player-btn': Button
    },
    props: [ 'value' ],
    computed: {
      row: function () {
        var self = this
        return self.rowText(self.value.area.row)
      }
    }
  }
</script>
