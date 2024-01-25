
export default function AddPageLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
      <body>{children}</body>
  );
}
