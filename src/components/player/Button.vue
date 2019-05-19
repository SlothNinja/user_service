<template>
  <sn-user-btn 
    :color="player.color"
    :user="player.user"
    :size="size"
  >
  </sn-user-btn>
</template>

<script>
  import Button from '@/components/user/Button'

  export default {
    name: 'sn-player-btn',
    props: [ 'player', 'size' ],
    components: {
      'sn-user-btn': Button
    }
  }
</script>
