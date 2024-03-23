<template>
    <div
        class="resizer"
        @mousedown="onResizeMouseDown"
        @mouseup="onResizeMouseUp"
    ></div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';


export interface Movement {
    direction: 'left' | 'right',
    change: number
}

const emit = defineEmits<{
    (e: 'move', value: Movement): void
}>()

let mousePosition = 0;

const resizeEvent = (event: MouseEvent) => {
    const change = mousePosition === 0 ? 1 : Math.abs(mousePosition - event.x);

    if (event.x > mousePosition) {
        emit('move', {
            direction: 'right',
            change
        })
    }
    if (event.x < mousePosition) {
        emit('move', {
            direction: 'left',
            change
        })
    }

    mousePosition = event.x;
};

const onResizeMouseDown = (): void => {
    window.addEventListener('mousemove', resizeEvent)
}

const onResizeMouseUp = (): void => {
    console.log('removing')
    window.removeEventListener('mousemove', resizeEvent);
}

onMounted(() => {
    window.addEventListener('mouseup', () => {
        console.log('mouseup');
        onResizeMouseUp();
    }, true)
})
</script>
<style scoped>
.resizer {
    width: 18px;
    height: 100%;
    background: blue;
    cursor: col-resize;
    position: relative;
}
</style>