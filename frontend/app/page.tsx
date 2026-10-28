export default function Home() {
  return (
    <main className="min-h-screen bg-slate-50">
      <nav className="mx-auto flex max-w-6xl items-center justify-between px-6 py-5">
        <a
          href="#"
          className="text-2xl font-bold tracking-tight text-slate-900"
        >
          Todo<span className="text-indigo-600">Flow</span>
        </a>

        <ul className="hidden items-center gap-8 text-sm font-medium text-slate-600 md:flex">
          <li>
            <a href="#" className="transition hover:text-indigo-600">
              Home
            </a>
          </li>
          <li>
            <a href="#features" className="transition hover:text-indigo-600">
              Features
            </a>
          </li>
          <li>
            <a href="#pricing" className="transition hover:text-indigo-600">
              Pricing
            </a>
          </li>
        </ul>

        <div className="flex items-center gap-5">

        <a
          href="#login"
          className="rounded-lg bg-indigo-600 px-4 py-2 text-sm font-semibold text-white shadow-sm transition hover:bg-indigo-700"
          >
          Sign up
        </a>
        <a
          href="#login"
          className="rounded-lg bg-indigo-600 px-4 py-2 text-sm font-semibold text-white shadow-sm transition hover:bg-indigo-700"
        >
          Sign in
        </a>
          </div>
        
      </nav>
    </main>
  );
}