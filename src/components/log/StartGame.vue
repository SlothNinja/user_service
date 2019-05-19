<template>
  <div>
    <div>
      Good luck:
    </div>
    <div
      v-for='(player, index) in value.players'
      :player='player'
      :key='index'
    >
      <sn-player-btn :player='player' size='small'></sn-player-btn>
      {{player.user.name}}
    </div>
    <div>
      Have fun.
    </div>
  </div>
</template>

<script>
  import Button from '@/components/player/Button'

  const _ = require('lodash')

  export default {
    name: 'sn-log-start-game-msg',
    props: [ 'value' ],
    components: {
      'sn-player-btn': Button
    },
    methods: {
      toString: function (list) {
        var size = _.size(list)
        switch (size) {
          case 0:
            return list
          case 1:
            return list
          case 2:
            return _.join(list, ' and ')
          default:
            return _.join(_.initial(list), ', ') + ', and ' + _.last(list)
        }
      }
    },
    computed: {
      ps: function () {
        return this.toString(_.map(this.value.players, 'user.name'))
      }
    }
  }
</script>
