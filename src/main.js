import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router/router'
import axios from 'axios'

Vue.config.productionTip = false

new Vue({
  data: {
    cu: null,
    idToken: '',
    snackbar: { open: false, message: '' },
    cp: {}
  },
  created () {
      var self = this
      self.fetchData()
  },
  methods: {
    fetchData () {
      var self = this
      axios.get('/current')
        .then(function (response) {
          self.cu = response.data.cu
        })
        .catch(function () {
          self.snackbar.message = 'Server Error.  Try again.'
          self.snackbar.open = true
          self.$router.push({ name: 'show', params: { id: self.$route.params.id}})
        })
    },
  },
  router,
  render: h => h(App),
}).$mount('#app')
