<template>
  <div>
    <tr v-for="feed in this.feeds" :key="feed.ID">
      <td>{{ feed.Url }}</td>
      <td>{{ feed.LastChecked }}</td>
    </tr>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: "FeedTable",
  data: function() {
    return {
      feeds: [],
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
    }
  }
}
</script>

<style scoped>

</style>