<template>
  <div v-if="isVisible" class="z-10 fixed bottom-32 right-32 animate-fade-in">
    <div role="alert" class="alert">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        class="stroke-info shrink-0 w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
        ></path>
      </svg>
      <span>{{ message }}</span>
    </div>
  </div>
</template>

<script lang="ts">
import { ref } from "vue";

export default {
    name: "AlertToast",
  props: {
    infoText: String,
    visible: Boolean,
  },
  
  setup(props: any) {
    const message = ref(props.infoText);
    const isVisible = ref(props.visible);

    return {
      message,
      isVisible,
    };
  },
  mounted() {
    setTimeout(() => {
      this.isVisible = true;
    }, 100);
  },
};
</script>

<style scoped>
@keyframes fade-in {
  from {
    opacity: 0;
    transform: translateY(20px) translateX(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0) translateX(-100px);
  }
}

.animate-fade-in {
  animation: fade-in 0.5s ease-out;
}
</style>
