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