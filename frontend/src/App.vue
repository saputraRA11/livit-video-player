<template>
  <div id="app">
    <h1 class="app-title">Video Search and Player</h1>
    <div class="app-container">
      <div class="content">
        <!-- Add ref attribute here to access SearchComponent as $refs.searchComponent -->
        <SearchComponent ref="searchComponent" @select-video="playVideo" @update-history="updateHistory" />
        <VideoPlayerComponent
          v-if="currentVideo"
          :videoUrl="currentVideo.url"
          :source="currentVideo.source"
        />
      </div>
      <!-- <div class="sidebar">
        <h2 class="history-title">Search History</h2>
        <ul class="history-list">
          <li v-for="keyword in searchHistory" :key="keyword" @click="loadFromHistory(keyword)">
            {{ keyword }}
          </li>
        </ul>
      </div> -->
    </div>
  </div>
</template>

<script>
import SearchComponent from "./components/SearchComponent.vue";
import VideoPlayerComponent from "./components/VideoPlayerComponent.vue";

export default {
  components: {
    SearchComponent,
    VideoPlayerComponent,
  },
  data() {
    return {
      currentVideo: null, // Holds the selected video data
      searchHistory: [] // Array to store search history
    };
  },
  methods: {
    playVideo(video) {
      this.currentVideo = video;
    },
    updateHistory(keyword) {
      // Add the keyword to history if it's not already present
      if (!this.searchHistory.includes(keyword)) {
        this.searchHistory.push(keyword);
        localStorage.setItem("searchHistory", JSON.stringify(this.searchHistory));
      }
    },
    loadFromHistory(keyword) {
      // Load the search from history by calling searchFromHistory on the SearchComponent
      this.$refs.searchComponent.searchFromHistory(keyword);
    }
  },
  mounted() {
    // Load search history from localStorage on component mount
    const savedHistory = JSON.parse(localStorage.getItem("searchHistory"));
    if (savedHistory) {
      this.searchHistory = savedHistory;
    }
  }
};
</script>

<style>
#app {
  text-align: center;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.app-title {
  font-size: 2em;
  font-weight: 600;
  color: #333;
  margin-bottom: 20px;
}

.app-container {
  display: flex;
  gap: 20px;
}

.content {
  flex: 3;
}

.sidebar {
  flex: 1;
  max-width: 250px;
  padding: 10px;
  border-left: 1px solid #ddd;
}

.history-title {
  font-size: 1.2em;
  font-weight: bold;
  color: #333;
  margin-bottom: 10px;
}

.history-list {
  list-style-type: none;
  padding: 0;
}

.history-list li {
  cursor: pointer;
  padding: 5px;
  border-radius: 5px;
  transition: background-color 0.2s;
}

.history-list li:hover {
  background-color: #f0f0f0;
}
</style>
