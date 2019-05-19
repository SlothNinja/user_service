<template>
  <div>
    <sn-toolbar app v-model="nav"></sn-toolbar>
    <sn-nav-drawer v-model="nav" app></sn-nav-drawer>
    <v-container grid-list-md >
      <v-layout row wrap>
        <v-flex xs6>
          <v-card height="31em">
            <v-card-title primary-title>
              <h3>New Game</h3>
            </v-card-title>
            <v-card-text>
              <v-form action="/got/game" method="post">
                <v-text-field
                  name="title"
                  label="Title"
                  auto="auto"
                  v-model="game.header.title"
                  id="title"
                >
                </v-text-field>
                <v-select
                  label="Number Players"
                  name="num-players"
                  auto="auto"
                  v-bind:items="npItems"
                  v-model="game.header.numPlayers"
                >
                </v-select> 
                <v-select 
                  id="two-thief-variant"
                  label="Two Thief Variant"
                  name="two-thief-variant"
                  auto="auto"
                  v-bind:items="optItems"
                  v-model="game.header.twoThief"
                >
                </v-select> 
                <v-text-field
                  label="Password"
                  name="password"
                  id="password"
                  v-model="game.header.password"
                  placeholder="Enter Password for Private Game"
                  type="password"
                  autocomplete="new-password"
                >
                </v-text-field>
                <v-btn @click="putData">Submit</v-btn>
              </v-form>
            </v-card-text>
          </v-card>
        </v-flex>
        <v-flex xs6>
          <v-card height="31em">
            <v-card-media :src="require('@/assets/got-box.jpg')" height="200px"></v-card-media>
            <v-card-text>
              <v-layout row>
                <v-flex xs5>Designer</v-flex>
                <v-flex>Adam E. Daulton</v-flex>
              </v-layout>
              <v-layout row>
                <v-flex xs5>Artists</v-flex> 
                <v-flex>Jeremy Montz</v-flex> 
              </v-layout>
              <v-layout row> 
                <v-flex xs5>Publisher</v-flex> 
                <v-flex><a href="http://www.thegamecrafter.com/">The Game Crafter, LLC</a></v-flex>
              </v-layout>
              <v-layout row>
                <v-flex xs5>Year Published</v-flex>
                <v-flex>2012</v-flex>
              </v-layout>
              <v-layout row> 
                <v-flex xs5>On-Line Developer</v-flex> 
                <v-flex>Jeff Huter</v-flex> 
              </v-layout> 
              <v-layout row> 
                <v-flex xs5>Permission Provided By</v-flex> 
                <v-flex>Adam E Daulton</v-flex> 
              </v-layout> 
              <v-layout row> 
                <v-flex xs5>Rules (pdf)</v-flex> 
                <v-flex><a href="/static/rules/got.pdf">Guild Of Thieves (English)</a></v-flex> 
              </v-layout> 
            </v-card-text>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
    <sn-footer></sn-footer>
  </div>
</template>

<script>
  import Toolbar from '@/components/Toolbar'
  import NavDrawer from '@/components/NavDrawer'
  import Footer from '@/components/Footer'

  export default {
    name: 'newGame',
    data () {
      return {
        game: {
          header: { title: '', id: 0, turn: 0, phase: 0, colorMaps: [], options: {} },
          state: { glog: [], jewels: {} }
        },
        auto: true,
        nav: false,
        npItems: [
          { text: '2', value: 2 },
          { text: '3', value: 3 },
          { text: '4', value: 4 }
        ],
        optItems: [
          { text: 'Yes', value: true },
          { text: 'No', value: false }
        ]
      }
    },
    components: {
      'sn-toolbar': Toolbar,
      'sn-nav-drawer': NavDrawer,
      'sn-footer': Footer
    },
    created () {
      var self = this
      self.fetchData()
    },
    watch: {
      '$route': 'fetchData'
    },
    computed: {
      cu: {
        get: function () {
          return this.$root.cu
        },
        set: function (value) {
          this.$root.cu = value
        }
      }
    },
    methods: {
      fetchData () {
        var self = this
        $.ajax({
          url: 'http://localhost:8081/got/game/new',
          dataType: 'json',
          success: function (data) {
            self.game.header = data.header
            self.game.state = data.state
            self.cu = data.cu
          }
        })
      },
      putData () {
        var self = this
        var game = {
          title: self.game.header.title,
          numPlayers: self.game.header.numPlayers,
          password: self.game.header.password,
          twoThief: self.game.header.twoThief
        }
        $.ajax({
          url: 'http://localhost:8081/got/game/new',
          dataType: 'json',
          data: JSON.stringify(game),
          contentType: 'application/json; charset=utf-8',
          type: 'PUT',
          success: function (data) {
            self.$router.push({name: 'index', params: { status: 'recruiting' }})
          }
        })
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  h1, h2, h3 {
    font-weight: normal;
  }
</style>
