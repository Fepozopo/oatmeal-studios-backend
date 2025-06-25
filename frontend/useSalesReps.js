// Utility to fetch all sales reps for the frontend
import { ref } from 'vue';

export function useSalesReps() {
    const salesReps = ref([]);
    const loading = ref(false);
    const error = ref(null);

    async function fetchSalesReps() {
        loading.value = true;
        error.value = null;
        try {
            const res = await fetch('/api/sales-reps');
            if (!res.ok) throw new Error('Failed to fetch sales reps');
            salesReps.value = await res.json();
        } catch (e) {
            error.value = e;
        } finally {
            loading.value = false;
        }
    }

    return { salesReps, loading, error, fetchSalesReps };
}
