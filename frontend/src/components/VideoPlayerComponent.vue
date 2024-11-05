<template>
  <div class="video-player-container">
    <!-- HTML5 Video Element for Direct Video Links -->
    <video
      v-if="source === 'direct'"
      controls
      :src="videoUrl"
      class="video-player"
    ></video>

    <!-- YouTube Embed -->
    <iframe
      v-else-if="source === 'youtube'"
      :src="youtubeEmbedUrl"
      width="100%"
      height="400"
      frameborder="0"
      allow="autoplay; encrypted-media"
      allowfullscreen
      class="iframe-player"
    ></iframe>

    <!-- TikTok Embed -->
    <iframe
      v-else-if="source === 'tiktok'"
      :src="tiktokEmbedUrl"
      width="100%"
      height="400"
      frameborder="0"
      allow="autoplay"
      allowfullscreen
      class="iframe-player"
    ></iframe>

    <!-- Message for unsupported video sources -->
    <p v-else>Unsupported video source.</p>
  </div>
</template>

<script>
export default {
  props: {
    videoUrl: {
      type: String,
      required: true,
    },
    source: {
      type: String,
      default: "direct", // Options: "direct", "youtube", "tiktok"
    },
  },
  computed: {
    youtubeEmbedUrl() {
      // Extract the video ID from various YouTube URL formats
      const videoId = this.extractYouTubeId(this.videoUrl);
      return videoId ? `https://www.youtube.com/embed/${videoId}` : '';
    },
    tiktokEmbedUrl() {
      // Convert TikTok URL to embed format
      return `https://www.tiktok.com/embed/${this.extractTikTokId()}`;
    },
  },
  methods: {
    extractYouTubeId(url) {
      const regex = /(?:youtube\.com\/(?:[^/]+\/.+\/|(?:v|embed|shorts)\/|.*[?&]v=)|youtu\.be\/)([^"&?/\s]{11})/;
      const match = url.match(regex);
      return match ? match[1] : null;
    },
    extractTikTokId() {
      const parts = this.videoUrl.split("/");
      return parts[parts.length - 1] || "";
    },
  },
};
</script>

<style scoped>
.video-player-container {
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
  background-color: #000;
  padding: 10px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.video-player {
  width: 100%;
  max-width: 100%;
  border-radius: 8px;
}

.iframe-player {
  width: 100%;
  border: none;
  border-radius: 8px;
}

@media (max-width: 768px) {
  .iframe-player,
  .video-player {
    height: 250px;
  }
}
</style>

