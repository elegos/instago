# InstaGo

## About
Easy image resizer for Instagram written in go

## License
Affero General Public License (AGPL) v3. See [LICENSE](LICENSE) file.

## Compilation
- `make`

## Usage
Copy the out folder's content somewhere you like, then drag'n'drop your image on the InstaGo.desktop file.
It will produce the image resized and converted for Instagram in the same directory of the original file, creating the new file `<original_file_name>.insta.jpg`. If it doesn't, you can use the application in console and see the errors, called by:

- `./instago /path/to/the/original/file`

If no image path is specified and `config.yml` is set to, all the images of where the binary is will be converted all together.

If images are too big they will be resized using the best match on the aspect ratiom as follows

Square (1:1)
Landscape (1.91:1)
Portrait (4:5)
Stories (9:16)
IGTV cover (1:1.55)

The software will though keep the original aspect ratio
