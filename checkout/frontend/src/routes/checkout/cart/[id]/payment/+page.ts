import { goto } from "$app/navigation";

async function delayNav(){
    await new Promise((resolve) => {
        setTimeout(resolve, 2000);
    })
    goto('/');
}