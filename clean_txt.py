import os
import glob
def clean_txt_files(directory):
    txt_files = glob.glob(os.path.join(directory, '*.txt'))
    for txt_file in txt_files:
        try:
            os.remove(txt_file)
            print(f"Removed: {txt_file}")
        except Exception as e:
            print(f"Error removing {txt_file}: {e}")
if __name__ == "__main__":
    directory = './data/'
    clean_txt_files(directory)
    directory = '.'
    clean_txt_files(directory)