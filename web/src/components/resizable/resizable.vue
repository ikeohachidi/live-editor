<template>
	<div class="resizable" ref="wrapperEl">
		<template
			v-for="(row, index) in resizableRows"
			:key="`row-${index}`"
			class="row"
		>
			<component
				:is="row"
				ref="resizableRow"
                :style="{ height: `${boxSize}px`}"
			></component>
			<Resizer
				v-if="index < (resizableRows.length - 1)"
				:before-el="resizableRow[index]?.$el"
				:after-el="resizableRow[index + 1]?.$el"
				axis="x"
			></Resizer>
		</template>
	</div>
</template>

<script setup lang="ts">
import {
	computed,
	useSlots,
	onMounted,
	ref,
	VNode
} from 'vue';
import Resizer from '../resizer/resizer.vue'

const slots = useSlots();

const resizableRows = ref<VNode[]>([]);
const resizableRow = ref<{$el: HTMLElement}[]>([]);
const wrapperEl = ref<HTMLElement>();
const boxSize = computed(() => {
    if (!wrapperEl.value) {
        return 0;
    }
    return wrapperEl.value?.getBoundingClientRect().height / resizableRows.value.length
})

onMounted(() => {
	if(slots && slots.default) {
		resizableRows.value = slots.default().filter((s) => {
			return ((s.type as unknown) as { name: string}).name === 'ResizableRow'
		});
	}
});
</script>

<style scoped>
.resizable {
	display: flex;
	flex-direction: column;
	height: 100vh;
}

.row {
	flex-grow: 1;
}
</style>