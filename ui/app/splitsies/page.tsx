export default function Page() {
    return (
        <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
            <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
                <div className="flex flex-col gap-4">
                    <button className="px-6 py-3 text-lg font-semibold bg-blue-600 text-white rounded-lg hover:bg-blue-700">
                        Add Item
                    </button>
                    <label className="px-6 py-3 text-lg font-semibold bg-blue-600 text-white rounded-lg hover:bg-blue-700 cursor-pointer">
                        Import Receipt
                        <input
                            type="file"
                            accept="image/*"
                            className="hidden"
                        />
                    </label>
                </div>
            </main>
        </div>
    )
}