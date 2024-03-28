<template>
	<div>
		<Resizable>
			<ResizableRow>
				<ResizableBox>
					<CodeEditor
						v-if="!isLoading"
						:sessionId="sessionId"
						v-model:content="htmlContent"
						language="html"
					/>
				</ResizableBox>
				<ResizableBox>
					<CodeEditor
						v-if="!isLoading"
						:sessionId="sessionId"
						v-model:content="cssContent"
						v-model:language="styleLanguage"
						:languageList="styleLanguges"
					/>
				</ResizableBox>
				<ResizableBox>
					<CodeEditor
						v-if="!isLoading"
						:sessionId="sessionId"
						v-model:content="jsContent"
						language="js"
					/>
				</ResizableBox>
			</ResizableRow>
			<ResizableRow>
				<ResizableBox>
					<div class="live-preview-frame">
						<iframe
							:id="String(sessionId)"
							class="live-preview"
							:src="`http://localhost:8000/session/${sessionId}`"
						></iframe>
					</div>
				</ResizableBox>
			</ResizableRow>
		</Resizable>
	</div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import Resizable from './components/resizable/resizable.vue';
import ResizableRow from './components/resizable-row/resizable-row.vue';
import ResizableBox from './components/resizable-box/resizable-box.vue';
import { type EditorLanguage } from './components/editor/editor.vue';
import CodeEditor from './components/editor/code-editor/code-editor.vue';

const htmlContent = ref('');
const cssContent = ref('');
const jsContent = ref('');

const styleLanguage = ref<EditorLanguage>('css');
const styleLanguges: {label: string, value: EditorLanguage}[] = [
	{
		label: 'Sass',
		value: 'sass'
	},
	{
		label: 'Css',
		value: 'css'
	}
];

const sessionId = 1;
const isLoading = ref(false);

const getSessionContent = async (): Promise<void> => {
	isLoading.value = true;

	try {
		const res = await fetch(`http://localhost:8000/content/${sessionId}`, {
			method: 'get',
			mode: 'cors',
		});
		const data = await res.json();
		htmlContent.value = data.html;
		cssContent.value = data.css;
		jsContent.value = data.js;

		isLoading.value = false;
	} catch(e) {
		console.error(e);
	}
}

onMounted(async() => {
	try {
		getSessionContent();
		await fetch(`http://localhost:8000/session/${sessionId}`, {
			method: 'post',
			mode: 'cors',
		});
		localStorage.setItem('editorId', String(sessionId));
	} catch(e) {
		console.error(e)
	}
})
</script>

<style scoped>
.wrapper {
	display: flex;
	justify-content: center;
	align-items: center;
	height: 100vh;
	width: 100vw;
}

.live-preview {
	border: none;
	border: 1px solid red;
	width: 100%;
	height: 100%;
	overflow: auto;
	position: absolute;
	top: 0;
	left: 0;
}

.live-preview-frame {
	position: relative;
	background-color: #fff;
	height: 100%;
}
</style>
