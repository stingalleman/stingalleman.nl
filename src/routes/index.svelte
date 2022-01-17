<!DOCTYPE html />
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<meta name="description" content="Sting Alleman" />
		<meta name="theme-color" content="#000000" />

		<link rel="preconnect" href="https://fonts.gstatic.com" />
		<link
			href="https://fonts.googleapis.com/css2?family=Nunito+Sans:wght@800&family=Nunito:wght@300;600&display=swap"
			rel="stylesheet"
		/>

		<title>Sting Alleman</title>
	</head>
	<body>
		<section>
			<div class="content">
				<header>
					<div class="heading">
						<h1>Sting</h1>
						<h1>Alleman</h1>
						<!-- <p>Subtitle.</p> -->
					</div>

					<nav>
						<a href="https://github.com/stingalleman">github</a>
						<a href="https://instagram.com/sting.alleman/">insta</a>
						<a href="https://discord.gg/SZ9Ept5CTM">discord</a>
						<a href="mailto:stingalleman@icloud.com">email</a>
					</nav>
				</header>

				<div class="window">
					<div class="top">
						<div class="dot" />
						<div class="dot" />
						<div class="dot" />
					</div>
					<pre>
            <p>hello!</p>
          </pre>
					<div class="bottom" />
				</div>
			</div>
		</section>
	</body>
</html>

<style lang="scss">
	:root {
		--bg: linear-gradient(45deg, #241c59, #020f2b);
	}

	*:not(span) {
		margin: 0px;
		padding: 0px;

		color: white;
		font-family: 'Nunito', sans-serif;
	}

	body {
		height: 100vh;
		width: 100vw;
		background-image: var(--bg);
	}

	section {
		position: absolute;
		width: 80vw;
		height: 80vh;

		left: 50%;
		top: 50%;
		transform: translate(-50%, -50%);

		background: #05011f;
		box-shadow: 0px 0px 100px rgba(0, 0, 0, 0.3);

		border-radius: 16pt;

		overflow: hidden;

		&::after {
			content: '';
			position: absolute;
			top: 0;

			width: 100%;
			height: 100%;

			background-image: linear-gradient(rgba(0, 0, 0, 0), rgba(0, 0, 0, 0.5));

			pointer-events: none;
		}
	}

	.content {
		width: 90%;
		height: 97%;

		margin-top: 3%;
		margin-left: 5%;
	}

	header {
		.heading {
			display: inline-block;

			h1 {
				margin-bottom: -0.6em;

				font-family: 'Nunito Sans', sans-serif;
				font-size: 3em;
			}

			// p {
			// 	margin-top: 2em;
			// 	margin-bottom: 4em;
			// }
		}

		nav {
			float: right;

			a {
				margin-right: 2em;
				text-decoration: none;

				&:hover {
					text-decoration: underline;
				}
			}
		}
	}

	.window {
		position: absolute;

		min-width: 700px;
		width: 60%;

		height: 75%;

		right: 5%;
		bottom: 0px;

		.top {
			display: flex;
			width: 100%;
			height: 3.2em;

			background: #2d2d2d;

			border-top-left-radius: 6px;
			border-top-right-radius: 6px;

			.dot {
				width: 12px;
				height: 12px;

				margin-top: 1.4em;
				margin-left: 8px;

				background-color: rgb(246, 195, 67);
				border-radius: 50%;

				&:first-child {
					margin-left: 1.4em;
					background-color: rgb(255, 23, 68);
				}

				&:last-child {
					background-color: rgb(92, 199, 42);
				}
			}
		}

		.bottom {
			display: flex;
			width: 100%;
			height: 1em;

			background: #2d2d2d;

			border-bottom-left-radius: 6px;
			border-bottom-right-radius: 6px;
		}

		pre {
			margin: 0px;

			max-height: calc(100% - 5em);
			overflow: auto scroll;
			background: #2d2d2d;

			white-space: pre-wrap;

			overflow-wrap: break-word;
		}
	}

	// @media (max-width: 970px) {
	// 	section {
	// 		top: 0px;
	// 		left: 0px;
	// 		transform: none;

	// 		width: 100vw;
	// 		height: 100vh;

	// 		border-radius: 0px;

	// 		.content {
	// 			width: 80%;
	// 			height: 90%;

	// 			margin: 5% 10%;
	// 		}
	// 	}

	// 	header {
	// 		display: flex;

	// 		.heading {
	// 			width: 100%;

	// 			font-size: 10pt;
	// 		}

	// 		nav {
	// 			display: flex;
	// 			flex-direction: column;

	// 			margin-top: 0.6em;

	// 			text-align: right;

	// 			a {
	// 				margin-right: 0px;
	// 			}
	// 		}
	// 	}

	// 	.window {
	// 		width: 90%;
	// 		min-width: auto;

	// 		margin-left: 5%;

	// 		font-size: 8pt;

	// 		pre {
	// 			border-bottom-left-radius: 6px;
	// 			border-bottom-right-radius: 6px;
	// 		}
	// 	}
	// }
</style>
