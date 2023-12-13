<template>
    <div role="alert" class="alert">
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
    <span>{{ message }}</span>
    </div>
</template>

<script lang="ts">
import { ref } from 'vue';

export default {
    Props: {
        content: "",
    },
    setup(props: any){
        const message = ref(props.content)
        return {
            message
        }
    },
}
</script>