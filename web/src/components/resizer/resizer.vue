<template>
    <div
        :class="[axis, 'resizer']"
        @mousedown="onResizeMouseDown"
        @mouseup="onResizeMouseUp"
    ></div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';


export interface Movement {
    direction: 'left' | 'right' | 'up' | 'down',
    change: number,
}

const props = withDefaults(defineProps<{
    axis: 'y' | 'x',
    beforeEl?: HTMLElement,
    afterEl?: HTMLElement
}>(), {
    axis: 'y'
});

const emit = defineEmits<{
    (e: 'move', value: Movement): void
}>()

let mousePosition = 0;

const onResizeMoveX = (movement: Movement): void => {
    const box = props.beforeEl;
    const afterBox = props.afterEl;
    if (!box) return;

    const boxWidth = parseFloat(box.style.width.replace('px', ''))
    const afterBoxWidth = afterBox ? parseFloat(afterBox.style.width.replace('px', '')) : 0

    if (movement.direction === 'right') {
        // increase the width of the box before the resizer
        // this box before the resizer is going to be at this index
        box.style.width = `${boxWidth+movement.change}px`;
        if (afterBox) {
            afterBox.style.width = `${afterBoxWidth-movement.change}px`;
        }
    }

    if (movement.direction === 'left') {
        box.style.width = `${boxWidth-movement.change}px`;
        if (afterBox) {
            afterBox.style.width = `${afterBoxWidth+movement.change}px`;
        }
    }
}

const onResizeMoveY = (movement: Movement): void => {
    const box = props.beforeEl;
    const afterBox = props.afterEl;
    if (!box) return;

    const boxHeight = parseFloat(box.style.height.replace('px', ''))
    const afterBoxHeight = afterBox ? parseFloat(afterBox.style.height.replace('px', '')) : 0

    if (movement.direction === 'up') {
        // increase the height of the box before the resizer
        // this box before the resizer is going to be at this index
        box.style.height = `${boxHeight+movement.change}px`;
        if (afterBox) {
            afterBox.style.height = `${afterBoxHeight-movement.change}px`;
        }
    }

    if (movement.direction === 'down') {
        box.style.height = `${boxHeight-movement.change}px`;
        if (afterBox) {
            afterBox.style.height = `${afterBoxHeight+movement.change}px`;
        }
    }
}

const resizeEvent = (event: MouseEvent) => {
    if (props.axis === 'x') {
        const change = mousePosition === 0 ? 1 : Math.abs(mousePosition - event.y);

        onResizeMoveY({
            direction: event.y < mousePosition ? 'down' : 'up',
            change
        });

        mousePosition = event.y;
    } else {
        const change = mousePosition === 0 ? 1 : Math.abs(mousePosition - event.x);

        onResizeMoveX({
            direction: event.x < mousePosition ? 'left' : 'right',
            change
        })

        mousePosition = event.x;
    }
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
    })
})
</script>
<style scoped>
.resizer {
    width: 18px;
    height: 100%;
    background-color: transparent;
    cursor: col-resize;
    position: relative;

    &.x {
        width: 100%;
        height: 18px;
        cursor: row-resize;
    }
}
</style>