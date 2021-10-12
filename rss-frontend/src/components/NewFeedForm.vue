<template>
  <tr>
    <td class="border border-green-600"><input type="text" v-model="newUrl"></td>
    <td class="border border-green-600"></td>
    <td class="border border-green-600">
      <button class="bg-green-400" v-on:click="submitUrl">Add</button>
    </td>
  </tr>
</template>

<script>
import axios from "axios";

export default {
  name: "NewFeedForm",
  data: function() {
    return {
      newUrl: "",
    }
  },
  methods: {
    submitUrl() {
      const url = 'http://localhost:8080/api/feed/new'
      axios.post(url, {
        Url: this.newUrl,
      })
          .then((response) => {
            console.log(response)
            this.newUrl = ""
            this.$emit('invalidate', 'new')
          })
    }
  }
}
</script>

<style scoped>

</style>