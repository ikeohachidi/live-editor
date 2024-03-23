<template>
    <div class="resizable-row" ref="wrapper">
        <template
            v-for="(box, index) in resizableBoxes"
            :key="`box-${index}`"
        >
            <component
                :is="box"
                ref="resizableBox"
                :style="{ width: `${boxSize}px`}"
            ></component>
            <Resizer
                v-if="index < (resizableBoxes.length - 1)"
                @move="onResizeMove($event, index)"
            ></Resizer>
        </template>
    </div>
</template>

<script lang="ts">
export default {
    name: 'ResizableRow'
}
</script>

<script setup lang="ts">
import {
    computed,
	useSlots,
	onMounted,
	ref,
	VNode,
} from 'vue';
import Resizer, { type Movement } from '../resizer/resizer.vue'

const slots = useSlots();

const resizableBoxes = ref<VNode[]>([]);
const resizableBox = ref<{$el: HTMLElement}[]>([]);
const boxSize = computed(() => {
    if (!wrapper.value) {
        return 0;
    }
    return wrapper.value?.getBoundingClientRect().width / resizableBoxes.value.length
})
const wrapper = ref<HTMLElement>()

const onResizeMove = (movement: Movement, index: number): void => {
    const box = resizableBox.value[index];
    const afterBox = resizableBox.value[index + 1];
    if (!box) return;

    const boxWidth = parseFloat(box.$el.style.width.replace('px', ''))
    const afterBoxWidth = parseFloat(afterBox.$el.style.width.replace('px', ''))

    console.log('change', movement.change, box.$el.style.width)
    if (movement.direction === 'right') {
        // increase the width of the box before the resizer
        // this box before the resizer is going to be at this index
        box.$el.style.width = `${boxWidth+movement.change}px`;
        if (afterBox) {
            afterBox.$el.style.width = `${afterBoxWidth-movement.change}px`;
        }
    }

    if (movement.direction === 'left') {
        box.$el.style.width = `${boxWidth-movement.change}px`;
        if (afterBox) {
            afterBox.$el.style.width = `${afterBoxWidth+movement.change}px`;
        }
    }
}

onMounted(() => {
	if(slots && slots.default) {
		resizableBoxes.value = slots.default().filter((s) => {
			return ((s.type as unknown) as { name: string}).name === 'ResizableBox'
		});
	}
});
</script>

<style scoped>
.resizable-row {
    padding: 0;
    margin: 0;
    border-bottom: 4px solid red;
    display: flex;
    flex-wrap: nowrap;
}

.resizable-row:deep( > div) {
    min-height: 300px;
}

.box {
    display: flex;
    flex-grow: 1;
}
</style>