<template>
    <div class="resizable-row" ref="wrapperEl">
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
                :before-el="resizableBox[index]?.$el"
                :after-el="resizableBox[index + 1]?.$el"
                axis="y"
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
import Resizer from '../resizer/resizer.vue'

const slots = useSlots();

const resizableBoxes = ref<VNode[]>([]);
const resizableBox = ref<{$el: HTMLElement}[]>([]);
const wrapperEl = ref<HTMLElement>()
const boxSize = computed(() => {
    if (!wrapperEl.value) {
        return 0;
    }
    return wrapperEl.value?.getBoundingClientRect().width / resizableBoxes.value.length
})

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
    display: flex;
    flex-wrap: nowrap;
    overflow: hidden;
}
</style>