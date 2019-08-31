<template>
  <v-flex xs4 class="text-xs-center">
    <v-tooltip bottom color='green'>
      <v-btn
        slot='activator'
        icon
        :disabled="!canReset"
              @click.native="$emit('action', { action: 'reset' })"
      >
        <v-icon>clear</v-icon>
      </v-btn>
      <span>Reset</span>
    </v-tooltip>
    <v-tooltip bottom color='green'>
      <v-btn
        slot='activator'
        icon
        :disabled="!canUndo"
        @click.native="$emit('action', { action: 'undo' })"
      >
        <v-icon>undo</v-icon>
      </v-btn>
      <span>Undo</span>
    </v-tooltip>
    <v-tooltip bottom color='green'>
      <v-btn
        slot='activator'
        icon
        :disabled="!canRedo"
        @click.native="$emit('action', { action: 'redo' })"
      >
        <v-icon>redo</v-icon>
      </v-btn>
      <span>Redo</span>
    </v-tooltip>
    <v-tooltip bottom color='green'>
      <v-btn
        slot='activator'
        icon
        :disabled="!canFinish"
        @click.native="$emit('action', { action : 'finish' })"
      >
        <v-icon>done</v-icon>
      </v-btn>
      <span>Finish</span>
    </v-tooltip>
  </v-flex>
</template>

<script>
  import UserButton from '@/components/user/Button'
  import Player from '@/components/mixins/Player'

  var _ = require('lodash')

  export default {
    name: 'sn-controlbar',
    mixins: [ Player ],
    components: {
      'sn-user-btn': UserButton
    },
    props: [ 'value' ],
    computed: {
      cu: {
        get: function () {
          return this.$root.cu
        },
        set: function (value) {
          this.$root.cu = value
        }
      },
      cp: {
        get: function () {
          return this.$root.cp
        },
        set: function (value) {
          this.$root.cp = value
        }
      },
      canUndo: function () {
        var self = this
        return (self.isCPorAdmin(self.value.header, self.value.state.players, self.cu) &&
          (self.value.undoStack.currentCount > self.value.undoStack.commitCount))
      },
      canRedo: function () {
        var self = this
        return (self.isCPorAdmin(self.value.header, self.value.state.players, self.cu) &&
          (self.value.undoStack.currentCount < self.value.undoStack.updateCount))
      },
      canReset: function () {
        var self = this
        return self.isCPorAdmin(self.value.header, self.value.state.players, self.cu)
      },
      canFinish: function () {
        var self = this
        return self.isCPorAdmin(self.value.header, self.value.state.players, self.cu) ? (_.get(self.cp, 'performedAction', true)) : false
      }
    }
  }
</script>
