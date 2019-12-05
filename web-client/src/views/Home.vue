<template>
  <div>
    <Nav />
    <LoggedOut
      v-if="state == 'lo'"
    />
    <Main
      v-else
    />
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue';
import Nav from '@/components/Nav.vue';
import LoggedOut from '@/components/LoggedOut.vue';
import Main from '@/components/Main.vue';

export default {
  name: 'home',
  components: {
    Nav,
    LoggedOut,
    Main,
  },
  data() {
    return {
      state: 'li',
    };
  },
};
</script>
