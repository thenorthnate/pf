<template>
  <div>
    <section class="section">
      <div class="container">
        <p class="title has-text-centered">Good evening Nate.</p>
        <nav class="level">
          <div class="level-item has-text-centered">
            <a class="button is-warning">
              <span class="icon is-large">
                <font-awesome-icon icon="dollar-sign" />
              </span>
              <span>Create Transaction</span>
            </a>
          </div>
          <div class="level-item has-text-centered">
            <a class="button is-primary">
              <span class="icon is-large">
                <font-awesome-icon icon="dollar-sign" />
              </span>
              <span>View Accounts</span>
            </a>
          </div>
          <div class="level-item has-text-centered">
            <a class="button is-danger">
              <span class="icon is-large">
                <font-awesome-icon icon="dollar-sign" />
              </span>
              <span>View Transactions</span>
            </a>
          </div>
          <div class="level-item has-text-centered">
            <a class="button is-link">
              <span class="icon is-large">
                <font-awesome-icon icon="dollar-sign" />
              </span>
              <span>View Plan</span>
            </a>
          </div>
        </nav>
        {{ message }}
        <div class="columns">
          <div class="column is-3">
            <div>
              <a class="button is-warning">
                <span class="icon is-large">
                  <font-awesome-icon icon="dollar-sign" />
                </span>
                <span>Create Transaction</span>
              </a>
            </div>
            <br>
            <a class="button is-primary">
              <span class="icon is-large">
                <font-awesome-icon icon="dollar-sign" />
              </span>
              <span>View Accounts</span>
            </a>
            <br>
            <a class="button is-danger">
              <span class="icon is-large">
                <font-awesome-icon icon="dollar-sign" />
              </span>
              <span>View Transactions</span>
            </a>
            <br>
            <a class="button is-link">
              <span class="icon is-large">
                <font-awesome-icon icon="dollar-sign" />
              </span>
              <span>View Plan</span>
            </a>

          </div>
          <div class="column">
            <div class="field">
              <p class="control has-icons-left has-icons-right">
                <input class="input" type="number" placeholder="Amount">
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="dollar-sign" />
                </span>
                <span class="icon is-small is-right">
                  <font-awesome-icon icon="check" />
                </span>
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import vNav from './v-nav';

export default {
  name: 'Main',
  components: {
    vNav,
  },
  data() {
    return {
      message: 'Hello, world!',
    };
  },
};
</script>
