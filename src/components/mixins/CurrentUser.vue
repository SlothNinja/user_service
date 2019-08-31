<script>
  export default {
    computed: {
      cu: {
        get: function () {
          return this.$root.cu
        },
        set: function (value) {
          this.$root.cu = value
        }
      },
      cuLoading: {
        get: function () {
          return this.$root.cuLoading
        },
        set: function (value) {
          this.$root.cuLoading = value
        }
      }
    }
  }
</script>
