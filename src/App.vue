<template>
	<div>
		<Resizable>
			<ResizableRow>
				<ResizableBox>
					Lorem ipsum dolor, sit amet consectetur adipisicing elit. Cumque animi in iure repellendus pariatur recusandae dolor perspiciatis. Minus necessitatibus natus corporis? Accusantium illum, cumque unde aliquid distinctio libero voluptatum exercitationem?
				</ResizableBox>
				<ResizableBox>
					Lorem ipsum dolor, sit amet consectetur adipisicing elit. Cumque animi in iure repellendus pariatur recusandae dolor perspiciatis. Minus necessitatibus natus corporis? Accusantium illum, cumque unde aliquid distinctio libero voluptatum exercitationem?
				</ResizableBox>
				<ResizableBox>
					Lorem ipsum dolor, sit amet consectetur adipisicing elit. Cumque animi in iure repellendus pariatur recusandae dolor perspiciatis. Minus necessitatibus natus corporis? Accusantium illum, cumque unde aliquid distinctio libero voluptatum exercitationem?
				</ResizableBox>
			</ResizableRow>
			<ResizableRow>
				<ResizableBox>
					Lorem ipsum dolor, sit amet consectetur adipisicing elit. Cumque animi in iure repellendus pariatur recusandae dolor perspiciatis. Minus necessitatibus natus corporis? Accusantium illum, cumque unde aliquid distinctio libero voluptatum exercitationem?
					<iframe src="https://ikeoha.com" frameborder="0"></iframe>
				</ResizableBox>
			</ResizableRow>
		</Resizable>
	</div>
</template>

<script setup lang="ts">
import Resizable from './components/resizable/resizable.vue';
import ResizableRow from './components/resizable-row/resizable-row.vue';
import ResizableBox from './components/resizable-box/resizable-box.vue';
</script>

<style scoped>
.wrapper {
	border: 5px solid blue;
	display: flex;
	justify-content: center;
	align-items: center;
	height: 100vh;
	width: 100vw;
}
</style>
