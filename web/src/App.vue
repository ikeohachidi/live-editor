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
						v-model:language="scriptLanguage"
						:languageList="scriptLanguages"
					/>
				</ResizableBox>
			</ResizableRow>
			<ResizableRow>
				<ResizableBox>
					<div class="live-preview-frame">
						<iframe
							v-if="sessionId"
							:id="String(sessionId)"
							class="live-preview"
							:src="sessionURL"
						></iframe>
					</div>
				</ResizableBox>
			</ResizableRow>
		</Resizable>

		<div class="links">
			<span>
				<a href="https://github.com/ikeohachidi/live-editor" target="_blank">Github</a>
			</span>
		</div>
	</div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import Resizable from './components/resizable/resizable.vue';
import ResizableRow from './components/resizable-row/resizable-row.vue';
import ResizableBox from './components/resizable-box/resizable-box.vue';
import { type EditorLanguage } from './components/editor/editor.vue';
import CodeEditor from './components/editor/code-editor/code-editor.vue';
import { fetchSession, createNewSession } from './http';
import { sessionIdStorageKey } from './http/constants';

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

const scriptLanguage = ref<EditorLanguage>('js');
const scriptLanguages: {label: string, value: EditorLanguage}[] = [
	{
		label: 'Js',
		value: 'js'
	},
	{
		label: 'Ts',
		value: 'ts'
	}
];

let sessionId = ref('');
const isLoading = ref(false);

const sessionURL = computed(() => {
    return `${import.meta.env.VITE_API}/session/${sessionId.value}`;
});

onMounted(async() => {
	isLoading.value = true;
	sessionId.value = localStorage.getItem(sessionIdStorageKey) || '';

	try {
		if (sessionId.value) {
			const { status, data } = await fetchSession(sessionId.value)
			if (status === 410) {
				const newSession = await createNewSession();
				localStorage.setItem(sessionIdStorageKey, String(newSession.sessionId));
				sessionId.value = newSession.sessionId;

				return;
			}

			htmlContent.value = data.html;
			cssContent.value = data.css;
			jsContent.value = data.js;

			return;
		}

		const newSession = await createNewSession();
		localStorage.setItem(sessionIdStorageKey, String(newSession.sessionId));
		sessionId.value = newSession.sessionId;
	} catch(e) {
		// TODO: handle this properly
		console.error(e);
	} finally {
		isLoading.value = false;
	}
})
</script>

<style scoped>
.links {
	position: fixed;
	bottom: 0;
	right: 0;
	background-color: rgba(0, 0, 0, 0.5);
	padding: 10px;

	a {
		color: #fff;
	}
}

.live-preview {
	border: none;
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
