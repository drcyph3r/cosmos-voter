<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
      <poll-list />
      <SpH3>
        Polls
      </SpH3>
      <poll-form />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";
import PollForm from "../components/PollForm";
import PollList from "../components/PollList";

export default {
  components: { PollList, PollForm, ...sp },
};
</script>
