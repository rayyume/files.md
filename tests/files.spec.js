// const {test, expect} = require('@playwright/test');
//
// test('should load files from OPFS', async ({ page }) => {
//     await page.goto(baseUrl);
//
//     // Override the function AFTER page loads but BEFORE it's called
//     await page.evaluate(() => {
//         window.getRootDirHandle = async function() {
//             // Your mock code here
//             const opfsRoot = await navigator.storage.getDirectory();
//             const testDir = await opfsRoot.getDirectoryHandle('test-files', { create: true });
//
//             const testFiles = [
//                 { name: 'README.md', content: '# Test\nHello world' },
//                 { name: 'Notes.md', content: '# Notes\n**Bold text**' }
//             ];
//
//             for (const fileData of testFiles) {
//                 try {
//                     await testDir.getFileHandle(fileData.name);
//                 } catch (error) {
//                     const fileHandle = await testDir.getFileHandle(fileData.name, { create: true });
//                     const writable = await fileHandle.createWritable();
//                     await writable.write(fileData.content);
//                     await writable.close();
//                 }
//             }
//
//             return testDir;
//         };
//     });
//
//     // Now call whatever triggers getRootDirHandle
//     await page.click('#some-button');
// });