<template>
  <span>Claimed {{data.from.card.type}} card at {{row}}{{data.from.column}}.</span>
</template>

<script>
  import Text from '@/components/mixins/Text'

  export default {
    mixins: [ Text ],
    name: 'sn-log-claim-item-msg',
    props: [ 'data' ],
    computed: {
      row: function () {
        var self = this
        return self.rowText(self.data.from.row)
      }
    }
  }
</script>
