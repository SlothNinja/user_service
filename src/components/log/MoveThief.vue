<template>
  <span>Moved thief to {{data.to.card.type}} at {{row}}{{data.to.column}}.</span>
</template>

<script>
  import Text from '@/components/mixins/Text'

  export default {
    mixins: [ Text ],
    name: 'sn-log-move-thief-msg',
    props: [ 'data' ],
    computed: {
      row: function () {
        var self = this
        return self.rowText(self.data.to.row)
      }
    }
  }
</script>
