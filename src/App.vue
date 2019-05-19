<template>
  <v-app>
    <sn-toolbar v-model="nav" ></sn-toolbar>
    <sn-snackbar v-model="snackbar.open">
      <div class="text-xs-center">
        {{snackbar.message}}
      </div>
    </sn-snackbar>
    <sn-nav-drawer v-model="nav" app></sn-nav-drawer>
    <v-content>    
      <v-container>
        <router-view></router-view>
      </v-container>
    </v-content>
    <sn-footer></sn-footer>
  </v-app>
</template>

<script>
  import Toolbar from '@/components/Toolbar'
  import NavDrawer from '@/components/NavDrawer'
  import Footer from '@/components/Footer'
  import Snackbar from '@/components/Snackbar'

  export default {
    name: 'App',
    data () {
      return {
        u: { name: '', emailReminders: 'true', emailNotifications: 'true', gravType: 'monsterid' },
        auto: true,
        nav: false,
        idToken: '',
        loading: true
      }
    },
    components: {
      'sn-snackbar': Snackbar,
      'sn-toolbar': Toolbar,
      'sn-nav-drawer': NavDrawer,
      'sn-footer': Footer
    },
    computed: {
      snackbar: {
        get: function () {
          return this.$root.snackbar
        },
        set: function (value) {
          this.$root.snackbar = value
        }
      }
    }
  }
</script>
