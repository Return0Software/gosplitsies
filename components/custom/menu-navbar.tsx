import Link from "next/link"
import { Sheet, SheetTrigger, SheetContent, SheetDescription, SheetHeader, SheetTitle } from "@/components/ui/sheet"
import { Button } from "@/components/ui/button"
import { VisuallyHidden } from "@radix-ui/react-visually-hidden"
import { Input } from "@/components/ui/input"

type MenuNavbarProps = {
    title: string
}

export default function MenuNavbar({ title = "Go Splitsies Title" }: MenuNavbarProps) {
    return (
        <nav className="fixed top-0 z-10 bg-white shadow dark:shadow-dark inset-x-0">
            <div className="w-full mx-auto max-w-7xl">
                <div className="flex items-center h-14 px-4 md:px-6">
                    <Link href="#" className="mr-6 text-2xl font-semibold" prefetch={false}>
                        Go Splitsies
                    </Link>

                    <div className="flex-1 flex justify-center">
                        <Input
                            className="max-w-sm border-none bg-transparent text-2xl font-semibold focus:border focus:border-input focus:bg-background"
                            type="text"
                            placeholder="Enter title..."
                            value={title}
                        />
                    </div>

                    <Button variant="ghost" size="icon" className="rounded-full">
                        <UserIcon className="w-5 h-5" />
                        <span className="sr-only">Toggle user menu</span>
                    </Button>

                    <Sheet>
                        <SheetTrigger asChild>
                            <Button variant="outline" className="ml-2 lg:hidden">
                                <MenuIcon className="w-4 h-4" />
                                <span className="sr-only">Toggle navigation menu</span>
                            </Button>
                        </SheetTrigger>

                        <SheetContent side="bottom">
                            <VisuallyHidden>
                                <SheetHeader>
                                    <SheetTitle>Edit profile</SheetTitle>
                                    <SheetDescription>
                                        Make changes to your profile here. Click save when you're done.
                                    </SheetDescription>
                                </SheetHeader>
                            </VisuallyHidden>
                            <div className="flex flex-col w-48 p-2 text-sm">
                                <Button variant="ghost" className="w-full justify-start">
                                    Profile
                                </Button>
                                <Button variant="ghost" className="w-full justify-start">
                                    Billing
                                </Button>
                                <Button variant="ghost" className="w-full justify-start">
                                    Settings
                                </Button>
                                <Button variant="ghost" className="w-full justify-start">
                                    Support
                                </Button>
                                <Button variant="ghost" className="w-full justify-start">
                                    Sign out
                                </Button>
                            </div>
                        </SheetContent>
                    </Sheet>
                </div>
            </div>
        </nav>
    )
}

function MenuIcon(props: React.SVGProps<SVGSVGElement>) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <line x1="4" x2="20" y1="12" y2="12" />
            <line x1="4" x2="20" y1="6" y2="6" />
            <line x1="4" x2="20" y1="18" y2="18" />
        </svg>
    )
}


function UserIcon(props: React.SVGProps<SVGSVGElement>) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
            <circle cx="12" cy="7" r="4" />
        </svg>
    )
}