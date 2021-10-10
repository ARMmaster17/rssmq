<template>
  <div>
    <table class="bg-white border-collapse border border-green-800 table-fixed w-11/12">
      <thead>
        <th class="border border-green-600 w-1/2">URL</th>
        <th class="border border-green-600 w-1/4">Last Updated</th>
        <th class="border border-green-600 w-1/4">Actions</th>
      </thead>
      <tbody>
        <tr v-for="feed in this.feeds" :key="feed.ID">
          <td class="border border-green-600">{{ feed.Url }}</td>
          <td class="border border-green-600">{{ feed.LastChecked }}</td>
          <td class="border border-green-600">
            <button class="bg-red-700" v-on:click="deleteUrl(feed.ID)">Delete</button>
          </td>
        </tr>
        <tr>
          <td class="border border-green-600"><input type="text" v-model="newUrl"></td>
          <td class="border border-green-600"></td>
          <td class="border border-green-600">
            <button class="bg-green-400" v-on:click="submitUrl">Add</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: "FeedTable",
  data: function() {
    return {
      feeds: [],
      newUrl: "",
    }
  },
  mounted: function() {
    this.getFeeds()
  },
  methods: {
    getFeeds() {
      const url = 'http://localhost:8080/api/feeds'
      axios.get(url, {
        dataType: 'json',
        mode: 'no-cors'
      })
      .then((response) => {
        console.log(response.data)
        this.feeds = response.data
      })
      .catch(function(error){
        console.log(error)
      })
    },
    submitUrl() {
      const url = 'http://localhost:8080/api/feed/new'
      axios.post(url, {
        Url: this.newUrl,
      })
      .then((response) => {
        console.log(response)
        this.newUrl = ""
        this.getFeeds()
      })
    },
    deleteUrl(id) {
      const url = 'http://localhost:8080/api/feed/' + id + '/delete'
      axios.post(url, {})
      .then((response) => {
        console.log(response)
        this.getFeeds()
      })
    }
  }
}
</script>

<style scoped>

</style>