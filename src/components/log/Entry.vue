<template>
  <div class='pa-2 mb-2'  style='border:1px solid black'>
    <v-toolbar
      dense
      height='20'
      :color='color'
      :dark='dark'
      flat
    >
      <div class='ml-1'>
        Turn: {{turn}}
      </div>
    </v-toolbar>
    <div v-for="(entry, index) in log" :key="index">
      <sn-log-message :value="entry"></sn-log-message>
    </div>
    <v-divider></v-divider>
    <div class="caption">
      {{createdAt}}
    </div>
  </div>
</template>

<script>
  import Message from '@/components/log/Message'
  import Color from '@/components/mixins/Color'

  var _ = require('lodash')

  export default {
    mixins: [ Color ],
    name: 'sn-log-entry',
    props: [ 'log' ],
    components: {
      'sn-log-message': Message
    },
    computed: {
      turn: function () {
        return _.get(_.last(this.log), 'data.turn', 0)
      },
      color: function () {
        return _.get(_.last(this.log), 'data.player.color', 'none')
      },
      dark: function () {
        return !(this.color === 'yellow')
      },
      createdAt: function () {
        var d = _.get(_.last(this.log), 'data.createdAt', false)
        if (d) {
          return new Date(d).toString()
        }
      }
    }
  }
</script>

<style scoped lang="scss">
  ul {
      display: block;
      list-style-type: disc;
      margin-top: 0;
      margin-bottom: 0;
      margin-left: 0;
      margin-right: 0;
      padding-left: 40px;
  }
</style>
