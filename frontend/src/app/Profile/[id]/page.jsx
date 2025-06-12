export default function Profile({params}) {
    
    return (
       <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <h1 className="text-2xl font-bold">Profile Page {params.id}</h1>
      <p className="text-lg">This is the profile page content.</p>
    </div>
    );
  }