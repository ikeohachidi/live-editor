<template>
	<div class="resizable">
		<div
			v-for="(row, index) in resizableRows"
			:key="`row-${index}`"
			class="row"
		>
			<component
				:is="row"
			></component>
		</div>
	</div>
</template>

<script setup lang="ts">
import {
	useSlots,
	onMounted,
	ref,
	VNode
} from 'vue';
import ResizableRow from '../resizable-row/resizable-row.vue';

const slots = useSlots();

const resizableRows = ref<VNode[]>([]);

onMounted(() => {
	if(slots && slots.default) {
		resizableRows.value = slots.default().filter((s) => {
			return ((s.type as unknown) as { name: string}).name === 'ResizableRow'
		});
	}
});
</script>