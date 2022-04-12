const electron = require('electron');
const path = require('path');
const url = require('url');

const { app } = electron;
const { BrowserWindow } = electron;

// const { ipcMain } = require('electron');

let mainWindow;

function createWindow() {
	const startUrl = process.env.DEV
		? 'http://localhost:3000'
		: url.format({
			pathname: path.join(__dirname, '/../build/index.html'),
			protocol: 'file:',
			slashes: true,
		});
	mainWindow = new BrowserWindow({
		width: 1920,
		height: 1280,
		icon: path.join(__dirname, 'Logo.png'),
		webPreferences: {
			nodeIntegration: true,
			contextIsolation: false,
		},
	});

	mainWindow.loadURL(startUrl);
	process.env.DEV && mainWindow.webContents.openDevTools();

	mainWindow.on('closed', function () {
		mainWindow = null;
	});
}

app.on('ready', createWindow);

app.on('window-all-closed', () => {
	if (process.platform !== 'darwin') {
		app.quit();
	}
});
