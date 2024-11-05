<template>
    <div class="search-container">
      <div class="search-bar">
        <input v-model="keyword" placeholder="Search videos..." class="search-input" />
        <button @click="searchVideos" class="search-button">Search</button>
      </div>
  
      <!-- Video grid, shown if videos are available -->
      <div v-if="videos && videos.length" class="video-list">
        <div
            v-for="video in videos"
            :key="video.id"
            class="video-item"
            @click="selectVideo(video)"
        >
            <img :src="video.thumbnail" :alt="video.title" class="thumbnail" />
            <div class="video-details">
            <p class="title">{{ video.title }}</p>
            <p class="author">By {{ video.author }}</p>
            </div>
        </div>
        </div>

  
      <div v-else>
        <p>No videos found.</p>
      </div>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    data() {
      return {
        keyword: "",
        videos: [],
      };
    },
    methods: {
      async searchVideos() {
        if (!this.keyword) return;
  
        try {
          const response = await axios.get("http://localhost:8080/search", {
            params: { keyword: this.keyword },
          });
          this.videos = response.data;
          this.$emit("update-history", this.keyword); // Emit the keyword to save in history
        } catch (error) {
          console.error("Error fetching videos:", error);
          this.videos = [];
        }
      },
      selectVideo(video) {
        this.$emit("select-video", video); // Emit the selected video to the parent
      },
      searchFromHistory(keyword) {
        this.keyword = keyword;
        this.searchVideos();
      }
    },
  };
  </script>
  
  <style scoped>
  .video-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
  max-width: 600px;
  margin: 0 auto;
  padding: 10px;
}

.video-item {
  display: flex;
  align-items: center;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 8px;
  transition: background-color 0.2s;
  cursor: pointer;
}

.video-item:hover {
  background-color: #f9f9f9;
}

.thumbnail {
  width: 80px;
  height: 80px;
  border-radius: 5px;
  margin-right: 15px;
  object-fit: cover;
}

.video-details {
  display: flex;
  flex-direction: column;
}

.title {
  font-size: 16px;
  font-weight: bold;
  margin: 0;
  color: #333;
}

.author {
  font-size: 14px;
  color: #666;
  margin-top: 5px;
}

  </style>
  