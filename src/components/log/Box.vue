<template>
  <v-container fluid>
    <v-layout row>
      <v-flex>
        <v-card color="green" dark>
          <v-card-title>
            <span class="title">Log</span>
            <v-spacer></v-spacer>{{logs.length}} of {{count}}
          </v-card-title>
        </v-card>
      </v-flex>
    </v-layout>
    <v-layout row justify-center>
      <v-flex>
        <v-card>
          <v-card-text>
            <v-container style="border:2px solid black;height:600px;overflow:scroll">
              <sn-log-entry
                class='pt-2'
                v-for="(log, index) in logs"
                :key="index"
                :log="log"
              >
              </sn-log-entry>
            </v-container>
          </v-card-text>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  import Entry from '@/components/log/Entry'

  var _ = require('lodash')

  export default {
    data: function () {
      return {
        offset: 0,
        loaded: false,
        logs: []
      }
    },
    props: [ 'count', 'open' ],
    components: {
      'sn-log-entry': Entry
    },
    name: 'sn-game-log',
    created () {
      var self = this
      self.fetchData()
    },
    watch: {
      count: function (newCount, oldCount) {
        var self = this
        if (self.loaded) {
          self.log = []
          self.loaded = false
        }
      },
      open: function (oldValue, newValue) {
        var self = this
        if ((self.open) && (!self.loaded)) {
          self.fetchData()
        }
      }
    },
    methods: {
      fetchData: _.debounce(
        function () {
          var self = this
          $.ajax({
            url: 'http://localhost:8081/got/game/glog/' + self.$route.params.id + '/' + self.count + '/' + self.offset,
            dataType: 'json',
            success: function (data) {
              if (data.logs) {
                self.offset = data.offset
                self.loaded = true
                self.logs = data.logs.concat(self.logs)
              }
            }
          })
        },
        500
      )
    }
  }
</script>
