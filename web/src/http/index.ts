import { type EditorLanguage } from "../components/editor/editor.vue";
import { Session } from "../types";
import { debounce } from "../utils/debounce";

const API = import.meta.env.VITE_API;

export const updateSession = async(sessionId: string, data: { content: string, lang: EditorLanguage }): Promise<void> => {
	try {
		await fetch(`${API}/session/${sessionId}`, {
			method: 'put',
			mode: 'cors',
			body: JSON.stringify(data)
		});

		const iframeEl = document.getElementById(String(sessionId)) as HTMLIFrameElement;
		if (iframeEl) {
			iframeEl.src = iframeEl.src;
		}
	} catch(e) {
		throw new Error(e as string);
	}
}

export const runUpdate = debounce(1000, updateSession);

export const fetchSession = async (sessionId: string): Promise<{ status: number, data: Session }> => {
	try {
		const res = await fetch(`${API}/content/${sessionId}`, {
			method: 'get',
			mode: 'cors',
		});

		if (res.status === 410) {
			return { status: 410, data: { html: '', css: '', js: '' } };
		}

		const data = await res.json();
		return { status: res.status, data };
	} catch(e) {
		throw new Error(e as string);
	}
}

export const createNewSession = async (): Promise<{ sessionId: string }> => {
	try {
		const res = await fetch(`${API}/session`, {
			method: 'post',
			mode: 'cors',
		});

		const data = await res.json();
		return data;
	} catch(e) {
		throw new Error(e as string);
	}
}