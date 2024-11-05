<template>
  <div class="watched-tags">
    <h3>Watched Tags</h3>
    <div v-if="tags.length" class="tags">
      <span v-for="tag in tags" :key="tag" class="tag">{{ tag }}</span>
    </div>
  </div>
</template>

<script>
export default {
  props: ["watchedTags"],
  computed: {
    tags() {
      return this.watchedTags;
    },
  },
};
</script>

<style scoped>
.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
}

.tag {
  background-color: #007bff;
  color: white;
  padding: 5px 10px;
  border-radius: 5px;
  font-size: 12px;
}
</style>
