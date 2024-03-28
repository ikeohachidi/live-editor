import { type EditorLanguage } from "../components/editor/editor.vue";
import { debounce } from "../utils/debounce";

export const updateContent = async(sessionId: number = 1, data: { content: string, lang: EditorLanguage }): Promise<void> => {
	try {
		await fetch(`http://localhost:8000/session/${sessionId}`, {
			method: 'put',
			mode: 'cors',
			body: JSON.stringify(data)
		});

		const iframeEl = document.getElementById(String(sessionId)) as HTMLIFrameElement;
		if (iframeEl) {
			iframeEl.src = iframeEl.src;
		}
	} catch(e) {
		console.error(e)
	}
}

export const runUpdate = debounce(1000, updateContent)