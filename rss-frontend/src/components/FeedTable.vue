<template>
  <div>
    <table class="bg-white border-collapse border border-green-800 table-fixed w-11/12">
      <thead>
        <th class="border border-green-600 w-1/2">URL</th>
        <th class="border border-green-600 w-1/4">Last Updated</th>
        <th class="border border-green-600 w-1/4">Actions</th>
      </thead>
      <tbody>
        <template v-for="feed in this.feeds">
          <FeedItem :item="feed" :key="feed.ID" v-on:invalidate="getFeeds" />
        </template>
        <NewFeedForm v-on:invalidate="getFeeds" />
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios'
import FeedItem from "@/components/FeedItem";
import NewFeedForm from "@/components/NewFeedForm";
export default {
  name: "FeedTable",
  components: {NewFeedForm, FeedItem},
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
      const url = '/api/feeds'
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