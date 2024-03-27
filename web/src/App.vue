<template>
	<div>
		<Resizable>
			<ResizableRow>
				<ResizableBox>
					<editor v-if="!isLoading" language="html" v-model="htmlContent"></editor>
				</ResizableBox>
				<ResizableBox>
					<editor v-if="!isLoading" language="css" v-model="cssContent"></editor>
				</ResizableBox>
				<ResizableBox>
					<editor v-if="!isLoading" language="javascript" v-model="jsContent"></editor>
				</ResizableBox>
			</ResizableRow>
			<ResizableRow>
				<ResizableBox>
					<div class="live-preview-frame">
						<iframe
							:id="String(sessionId)"
							class="live-preview"
							:src="`http://localhost:8000/session/${sessionId}`"
							style="width: 100%; height: 100%"
						></iframe>
					</div>
				</ResizableBox>
			</ResizableRow>
		</Resizable>
	</div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import Resizable from './components/resizable/resizable.vue';
import ResizableRow from './components/resizable-row/resizable-row.vue';
import ResizableBox from './components/resizable-box/resizable-box.vue';
import Editor, { type EditorLanguage } from './components/editor/editor.vue';
import { debounce } from './utils/debounce';

const htmlContent = ref('');
watch(htmlContent, (content: string) => {
	runUpdate(content, 'html' as EditorLanguage)
})
const cssContent = ref('');
watch(cssContent, (content: string) => {
	runUpdate(content, 'css' as EditorLanguage)
})
const jsContent = ref('');
watch(jsContent, (content: string) => {
	runUpdate(content, 'js' as EditorLanguage)
})

const sessionId = 1;
const isLoading = ref(false);

const updateContent = async(content: string, lang: EditorLanguage): Promise<void> => {
	try {
		const sessionId = 1;
		await fetch(`http://localhost:8000/session/${sessionId}`, {
			method: 'put',
			mode: 'cors',
			body: JSON.stringify({
				lang,
				content,
			})
		});

		const iframeEl = document.getElementById(String(sessionId)) as HTMLIFrameElement;
		if (iframeEl) {
			iframeEl.src = iframeEl.src;
		}
	} catch(e) {
		console.error(e)
	}
}

const runUpdate = debounce(1000, updateContent)

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
}

.live-preview-frame {
	background-color: #fff;
	height: 100%;
}
</style>
