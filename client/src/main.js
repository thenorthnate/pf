// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import { library } from '@fortawesome/fontawesome-svg-core';
import {
  faCheck,
  faDollarSign,
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
// import { FontAwesomeIcon, FontAwesomeLayers } from '@fortawesome/vue-fontawesome';
import 'bulma/css/bulma.css';
import Vue from 'vue';
import App from './App';
import router from './router';

library.add(faCheck);
library.add(faDollarSign);
Vue.component('font-awesome-icon', FontAwesomeIcon);
// Vue.component('font-awesome-layers', FontAwesomeLayers);

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
});
