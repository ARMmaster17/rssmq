<template>
  <tr >
    <td class="border border-green-600">{{ item.Url }}</td>
    <td class="border border-green-600">{{ item.LastChecked }}</td>
    <td class="border border-green-600">
      <button class="bg-red-700" v-on:click="deleteUrl(item.ID)">Delete</button>
    </td>
  </tr>
</template>

<script>
import axios from "axios";

export default {
  name: "FeedItem",
  props: {
    item: {
      type: Object,
      required: true
    }
  },
  methods: {
    deleteUrl(id) {
      const url = 'http://localhost:8080/api/feed/' + id + '/delete'
      axios.post(url, {})
          .then((response) => {
            console.log(response)
            this.$emit('invalidate', 'delete')
          })
    }
  }
}
</script>

<style scoped>

</style>