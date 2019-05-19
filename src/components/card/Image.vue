<template>
  <div class='got-card' :class="show && kind ? kind : 'card-back'" ></div>
</template>

<script>
  export default {
    name: 'sn-card-image',
    props: [ 'kind', 'show' ]
  }
</script>

<style scoped lang="scss">

  .got-card {
    width:100px;
    height:136px;
    background-image:url('../../assets/cards.png');
    border-radius:10px;
    margin:2px;
  }

  .got-card.none {
    opacity:0;
    position:absolute;
    top:0px;
    left:0px;
  }

  #playable-card, #shufflable-card, #drawable-card, #claimable-card {
    position:absolute;
    top:0px;
    left:0px;
    pointer-events:none;
    .none {
      opacity:0;
    }
  }

  #playable-card, #shufflable-card, #claimable-card {
    z-index:1000;
  }

  #drawable-card {
    z-index:1001;
  }

  @mixin got-card($class, $left) {
    .got-card.#{$class} {
      background-position:$left 0%;
      background-size:cover;
    }
  }

  @include got-card(camel, 0%);
  @include got-card(carpet, 10%);
  @include got-card(coins, 20%);
  @include got-card(guard, 30%);
  @include got-card(jewels, 40%);
  @include got-card(lamp, 50%);
  @include got-card(start-camel, 60%);
  @include got-card(start-lamp, 70%);
  @include got-card(sword, 80%);
  @include got-card(turban, 90%);
  @include got-card(card-back, 100%);

</style>
