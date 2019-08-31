<template>
  <v-toolbar height="128" scroll-off-screen :scroll-threshold="10" clipped-left flat color="green" dark app >
    <v-layout row>
      <v-flex xs6>
        <v-layout row>
          <v-flex>
            <div style="height:4em" v-if="cuLoading">
              &nbsp;
            </div>
            <div style="height:4em" v-else>
              <div v-if="cu" class="font-weight-bold title">
                <sn-user-btn size="small" :user="cu" ></sn-user-btn>
                {{cu.name}}
                <v-tooltip bottom>
                  <template v-slot:activator="{ on }">
                    <v-btn @click="logout" icon color="green" v-on="on">
                      <v-icon>exit_to_app</v-icon>
                    </v-btn>
                  </template>
                  <span>Logout</span>
                </v-tooltip>
              </div>
              <div v-else>
                <v-btn href="/login" color="green">Login</v-btn>
              </div>
            </div>
          </v-flex>
        </v-layout>
        <v-layout row>
          <v-flex>
            <v-tooltip bottom>
              <template v-slot:activator="{ on }">
                <div>
                  <v-toolbar-side-icon v-on="on" @click.stop="$emit('input', !value)" ></v-toolbar-side-icon>
                </div>
              </template>
              <span>Menu</span>
            </v-tooltip>
          </v-flex>
        </v-layout>
      </v-flex>
      <v-spacer></v-spacer>
      <v-flex xs2 class='text-xs-right' color="black">
          <v-card :to="{ name: 'home' }" color="white" height='100'>
          <v-img max-height='100' contain :src="require('../assets/slothninja_logo_fullsize.png')" />
        </v-card>
      </v-flex>
    </v-layout>
  </v-toolbar>
</template>


<script>
  import UserButton from '@/components/user/Button'
  import CurrentUser from '@/components/mixins/CurrentUser'

  export default {
    mixins: [ CurrentUser ],
    name: 'sn-toolbar',
    components: {
      'sn-user-btn': UserButton
    },
    props: [ 'value' ],
    methods: {
      logout: function () {
        var self = this
        self.delete_cookie('sngsession')
        self.cu = null
        self.cuLoading = false
        self.$router.push({ name: 'home'})
      },
      delete_cookie: function (name) {
        document.cookie = name + '= ; domain = .slothninja.com ; expires = Thu, 01 Jan 1970 00:00:00 GMT'
      }
    }
  }
</script>

<style scoped lang="scss">
  img.logo {
    height: 100px;
    background: white;
    border-radius:10px;
  }
</style>
