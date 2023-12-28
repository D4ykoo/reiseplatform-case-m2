// import { error } from '@sveltejs/kit';
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params }) => {
	const id = Number(params.id);

	if (isNaN(id)){
		error(404, 'Not found');
	}
	
	const res = await fetch(`http://localhost:8071/c/get/${params.id}`);
	const data = await res.json()

	return {id: params.id, data: data}
	
};