<script lang="ts">
	import { page } from "$app/stores";
    import type { PageData } from "./$types";
    export let data: PageData;
</script>

<div class="pl-8 pr-8">
    <h1>Cart {data.id}</h1>
    <div class="flex flex-col">
        <p>id: {data.data.id}</p>
        <p>paid: {data.data.paid}</p>
        <p>payment method: {data.data.payment_method} </p>
        <p>offers: {data.data.offers} </p>
    </div>

    <a class="btn" href="{$page.url.pathname}/payment">Pay</a>
</div>